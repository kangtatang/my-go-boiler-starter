/**
 * Skip to Content Utility
 * Updates the "Skip to Content" link dynamically based on the current page.
 */
const SkipContent = {
    /**
     * Initializes the skip-to-content functionality.
     * @param {string} defaultContentId - The default content ID to use.
     */
    init(defaultContentId) {
      this.defaultContentId = defaultContentId || 'main-content';
      this.skipLink = document.querySelector('.skip-link');
  
      if (!this.skipLink) {
        console.error('Skip to Content link not found.');
        return;
      }
  
      // Update the skip link on page load and URL changes
      window.addEventListener('hashchange', () => this._updateSkipLink());
      window.addEventListener('DOMContentLoaded', () => this._updateSkipLink());
    },
  
    /**
     * Updates the skip-to-content link based on the current page.
     */
    _updateSkipLink() {
      const currentPage = window.location.hash.slice(1); // Get hash without '#'
      let contentId;
  
      switch (currentPage) {
        case 'home':
          contentId = 'home-content'; // Adjust this to match the home content ID
          break;
        case 'detail':
          contentId = 'detail-content'; // Adjust this to match the detail content ID
          break;
        default:
          contentId = this.defaultContentId; // Use default if no match
      }
  
      // Set the ID for the main content
      const mainContent = document.querySelector('main');
      if (mainContent) {
        mainContent.setAttribute('id', contentId);
      } else {
        console.warn('Main content not found. Make sure your main element exists.');
      }
  
      // Update the skip link href attribute
      this.skipLink.setAttribute('href', `#${contentId}`);
    },
  };
  
  export default SkipContent;
  